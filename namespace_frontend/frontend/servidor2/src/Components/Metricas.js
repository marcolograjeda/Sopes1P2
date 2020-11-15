import React, { Component } from 'react';
import Chart from "react-google-charts";
import axios from 'axios';


class Metricas extends Component {
    url = 'appserver:3000'
    constructor() {
        super();
        this.state = { departamentos: [], top3:[], custom: [], ultimo:{
            name: '',
            location: '',
            age: '',
            infectedtype: '',
            state: ''
        }, edades:[]}
    }
    componentDidMount() {
        this.getDataAxios()
    }

    async getDataAxios() {
        try {
            let response = await axios.get(this.url + "/departamentos")
            //console.log(response.data)
            let arreglo=[['Departamento', 'Infectados']];
            for (const iterator of response.data) {
                arreglo.push([iterator.location, iterator.cantidad]);
            }
            this.setState({custom:arreglo});
            if(response.data.length>3){
                let arr=[];
                for (let index = 0; index < 3; index++) {
                    const element = response.data[index];
                    arr.push(element);
                }
                this.setState({departamentos: response.data, top3: arr});
            }else{
                this.setState({ departamentos: response.data, top3:response.data });
            }

            //datos de ultimo caso agregado
            let last_response = await axios.get(this.url + "/last");
            //console.log(last_response.data.ultimo);
            this.setState({ultimo:last_response.data.ultimo});

            //datos de edades
            let edades_response = await axios.get(this.url + "/edades")
            let edades=[['Edad', 'Cantidad actual']];
            for (const iterator of edades_response.data) {
                if(iterator.lim===10){
                    edades.push(['0 a 10', iterator.cantidad]);
                }else if(iterator.lim===20){
                    edades.push(['11 a 20', iterator.cantidad]);
                }else if(iterator.lim===30){
                    edades.push(['21 a 30', iterator.cantidad]);
                }else if(iterator.lim===40){
                    edades.push(['31 a 40', iterator.cantidad]);
                }else if(iterator.lim===50){
                    edades.push(['41 a 50', iterator.cantidad]);
                }else if(iterator.lim===60){
                    edades.push(['51 a 60', iterator.cantidad]);
                }else if(iterator.lim===70){
                    edades.push(['61 a 70', iterator.cantidad]);
                }else if(iterator.lim===80){
                    edades.push(['71 a 80', iterator.cantidad]);
                }else if(iterator.lim===90){
                    edades.push(['81 a 90', iterator.cantidad]);
                }else if(iterator.lim===100){
                    edades.push(['91 a 100', iterator.cantidad]);
                }else if(iterator.lim===110){
                    edades.push(['101 a 110', iterator.cantidad]);
                }else if(iterator.lim===120){
                    edades.push(['111 en adelante', iterator.cantidad]);
                }
            }
            this.setState({edades:edades});
            
        } catch (error) {
            console.log('no se pudo obtener respuesta del servidor')
        }
    }

    handleClick = async () => {
        await this.getDataAxios();
      }


    render() {
        return (
            <div className="container" style={{ marginTop: '2%' }} id="metricas">
                <button className="btn btn-outline-primary float-right" style={{ marginBottom: '1%' }} type="button" onClick={this.handleClick}>Actualizar</button>
                <div className="row">
                    <div className="col-md-4">
                        {/*Top3*/}
                        <strong>Top 3  de departamentos con coronavirus:</strong>
                        <hr />
                        {
                            this.state.top3.map(function(dep, index){
                            return <div key={index}>{dep.location}: {dep.cantidad} casos <hr /></div>
                            })
                        }

                    </div>
                    <div className="col-md-8">
                        {/*Pie*/}
                        <Chart
                            width={'500px'}
                            height={'300px'}
                            chartType="PieChart"
                            loader={<div>Loading Chart</div>}
                            data={this.state.custom}
                            options={{
                                title: 'Todos los departamentos afectados',
                            }}
                            rootProps={{ 'data-testid': '1' }}
                        />
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-4">
                        <strong>Último caso agregado:</strong>
                        <hr />
                        <div>
                        name: { this.state.ultimo.name===''? <strong>NO TIENE CASOS REDIS</strong>: this.state.ultimo.name}
                        </div>
                        <hr />
                        <div>
                        Location: { this.state.ultimo.location===''? <strong>NO TIENE CASOS REDIS</strong>: this.state.ultimo.location}
                        </div>
                        <hr />
                        <div>
                        Age: { this.state.ultimo.age===''? <strong>NO TIENE CASOS REDIS</strong>: this.state.ultimo.age}
                        </div>
                        <hr />
                        <div>
                        infectedtype: { this.state.ultimo.infectedtype===''? <strong>NO TIENE CASOS REDIS</strong>: this.state.ultimo.infectedtype}
                        </div>
                        <hr />
                        <div>
                        state: { this.state.ultimo.state===''? <strong>NO TIENE CASOS REDIS</strong>: this.state.ultimo.state}
                        </div>
                    </div>
                    <div className="col-md-8">
                    <Chart
                        width={'500px'}
                        height={'300px'}
                        chartType="BarChart"
                        loader={<div>Loading Chart</div>}
                        data={this.state.edades}
                        options={{
                            title: 'Edades afectadas',
                            chartArea: { width: '50%' },
                            hAxis: {
                            title: 'Afectados',
                            minValue: 0,
                            },
                            vAxis: {
                            title: 'Edades',
                            },
                        }}
                        // For tests
                        rootProps={{ 'data-testid': '1' }}
                    />
                    </div>
                </div>
            </div>
        );
    }

}
export default Metricas;