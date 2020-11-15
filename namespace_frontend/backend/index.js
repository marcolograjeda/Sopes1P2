const { MongoClient } = require('mongodb')
const express = require('express')
const app = express();
const bodyParser = require('body-parser');
var redis = require('redis');
var client_redis = redis.createClient('6379', 'ec2-3-139-95-193.us-east-2.compute.amazonaws.com'); //creates a new client

client_redis.on('connect', function () {
    console.log('connected');
});

app.use(express.json({ limit: '50mb' }));
app.use(bodyParser.urlencoded({ extended: true }));


const uri = 'mongodb://ec2-3-139-95-193.us-east-2.compute.amazonaws.com:27017';
const client = new MongoClient(uri, { useNewUrlParser: true, useUnifiedTopology: true });

app.use((req, res, next) => {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Headers', 'Authorization, X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Allow-Request-Method');
    res.header('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, DELETE');
    res.header('Allow', 'GET, POST, OPTIONS, PUT, DELETE');
    next();
});

app.get('/', (req, res) => {
    res.send('<h1>Funciona!</h1>');
});

app.post('/insertmongo', async (req, res) => {
    const newCaso = {
        name: req.body.name,
        location: req.body.location,
        age: req.body.age,
        infectedtype: req.body.infectedtype,
        state: req.body.state
    };
    let result = { 'message': 'ok' };
    try {
        await client.connect();
        const database = client.db('test');
        const collection = database.collection('casos')
        collection.insertOne(newCaso);
    } catch (err) {
        console.log(err);
        result = { 'message': 'failed' };
    } finally {
        res.json(result);
    }
});

app.post('/insertredis', async (req, res) => {
    const newCaso = {
        name: req.body.name,
        location: req.body.location,
        age: req.body.age,
        infectedtype: req.body.infectedtype,
        state: req.body.state
    };
    client_redis.get('casos', function (err, reply) {
        //console.log(reply);
        let original = reply;
        if(original==null){
            original=[]
        }else{
            console.log('ya tiene elementos!');
            original=JSON.parse(original);
        }
        original.push(newCaso);
        original=JSON.stringify(original);
        //aqui le inserto el arreglo ya con el nuevo elemento
        client_redis.set("casos", original, function (err, reply) {
            console.log(reply);
        });
        
    });
    res.json({ 'message': 'ok' });
});

app.get('/last',(req, res)=>{
    client_redis.get('casos', function (err, reply) {
        //console.log(reply);
        if(reply==null){
            res.json({ultimo:{
                name: '',
                location: '',
                age: '',
                infectedtype: '',
                state: ''
            }});
        }else{
            let original=JSON.parse(reply);
            res.json({ultimo:original[original.length-1]});
        }
        
    });
});

app.get('/departamentos', async (req, res) => {
    MongoClient.connect(uri, { useNewUrlParser: true, useUnifiedTopology: true }, function (err, db) {
        if (err) throw err;
        var dbo = db.db("test");
        dbo.collection("casos").find({}).toArray(function (err, result) {
            if(err){
                res.status(500).json(err);
            }else{
                res.status(200).json(sortTop3(result));
            }
            db.close();
        });
    });

});

app.get('/edades',(req, res)=>{
    client_redis.get('casos', function (err, reply) {
        //console.log(reply);
        if(reply==null){
            res.status(200).json([]);
        }else{
            let original=JSON.parse(reply);
            res.status(200).json(cantidadAfectados(original));
        }
        
    });
});

function cantidadAfectados(arreglo){
    let res=[
        {
            lim:10,
            cantidad:0
        },
        {
            lim:20,
            cantidad:0
        },
        {
            lim:30,
            cantidad:0
        },
        {
            lim:40,
            cantidad:0
        },
        {
            lim:50,
            cantidad:0
        },
        {
            lim:60,
            cantidad:0
        },
        {
            lim:70,
            cantidad:0
        },
        {
            lim:80,
            cantidad:0
        },
        {
            lim:90,
            cantidad:0
        },
        {
            lim:100,
            cantidad:0
        },
        {
            lim:110,
            cantidad:0
        },
        {
            lim:120,
            cantidad:0
        }

    ];
    for (const iterator of arreglo) {
        if(iterator.age>=0 && iterator.age<=10){
            res.find(x => x.lim === 10).cantidad++;
        }else if(iterator.age>10 && iterator.age<=20){
            res.find(x => x.lim === 20).cantidad++;
        }else if(iterator.age>20 && iterator.age<=30){
            res.find(x => x.lim === 30).cantidad++;
        }else if(iterator.age>30 && iterator.age<=40){
            res.find(x => x.lim === 40).cantidad++;
        }else if(iterator.age>40 && iterator.age<=50){
            res.find(x => x.lim === 50).cantidad++;
        }else if(iterator.age>50 && iterator.age<=60){
            res.find(x => x.lim === 60).cantidad++;
        }else if(iterator.age>60 && iterator.age<=70){
            res.find(x => x.lim === 70).cantidad++;
        }else if(iterator.age>70 && iterator.age<=80){
            res.find(x => x.lim === 80).cantidad++;
        }else if(iterator.age>80 && iterator.age<=90){
            res.find(x => x.lim === 90).cantidad++;
        }else if(iterator.age>90 && iterator.age<=100){
            res.find(x => x.lim === 100).cantidad++;
        }else if(iterator.age>100 && iterator.age<=110){
            res.find(x => x.lim === 110).cantidad++;
        }else if(iterator.age>110){
            //cual quier otra mayor a 110
            res.find(x => x.lim === 120).cantidad++;
        }
    }
    return res;
}

function sortTop3(arreglo) {
    let heads = [];
    let res = [];
    //primero jalamos todos los location que hay
    for (const iterator of arreglo) {
        if (!heads.includes(iterator.location)) {
            heads.push(iterator.location);
            res.push({ location: iterator.location, cantidad: 1 })
        } else {
            res.find(x => x.location === iterator.location).cantidad++;
        }
    }
    res.sort((a, b) => b.cantidad - a.cantidad);
    return res;
}


app.listen(3000, () => {
    console.log('listening at port 3000');
});