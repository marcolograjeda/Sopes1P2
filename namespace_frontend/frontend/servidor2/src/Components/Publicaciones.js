import React, { Component } from 'react';
import axios from 'axios';


class Publicaciones extends Component {
  url_A = 'http://18.223.114.187'
  url_B = 'http://18.219.119.12'
  tieneA=true;
  tieneB=true;
  constructor() {
    super();
    this.state = { users: [], selected: true }
  }

  componentDidMount() {
    //fetch('http://18.223.114.187:3000/poetas')
    // .then(response => console.log(response.json()));
    //.then(json => this.setState({ users: json.data }));
    //let myobj={autor:"ronald2", nota:"nota de prueba2"}
    //this.setState({users:[myobj, myobj]})
    this.getDataAxios()

  }

  async getDataAxios() {
    let respuestaA;
    let respuestaB;
    try {
      respuestaA = await axios.get(this.url_A + ":3000/poetas")
      //console.log(response.data.data)
      //this.setState({ users: respuestaA.data.data })  
    } catch (error) {
      this.tieneA=false;
    }

    try {
      respuestaB = await axios.get(this.url_B + ":3000/poetas")
      //console.log(response.data.data)
      //this.setState({ users: respuestaA.data.data })  
    } catch (error) {
      this.tieneB=false;
    }

    if(this.tieneA && this.tieneB){
      //le damos a A
      this.setState({ users: respuestaA.data.data })  
    }else{
      if(this.tieneA){
        //nos quedamos para iniciar con A
        this.setState({ users: respuestaA.data.data })  
      }else if(this.tieneB){
        //nos quedamos con B
        this.setState({ users: respuestaB.data.data })  
        this.setState({ selected: false });
      }
    }

  }

  handleClick = async () => {
    if (this.state.selected) {
      //servidorA
      if(this.tieneB){
        const response = await axios.get(this.url_B + ":3000/poetas")
        //console.log(response.data.data)
        this.setState({ users: response.data.data })
        this.setState({ selected: false });
      }else{
        alert('No se puede obtener respuesta del servidor1B');
      }
    } else {
      //servidorB
      if(this.tieneA){
        const response = await axios.get(this.url_A + ":3000/poetas")
        //console.log(response.data.data)
        this.setState({ users: response.data.data })
        this.setState({ selected: true });
      }else{
        alert('No se puede obtener respuesta del servidor1A');
      }
    }
  }
  render() {
    return (
      <section className="page-section " id="publicaciones">


        <div className="container">
          <div className="row">
            {
              this.state.selected ? <h1>Publicaciones Servidor1A</h1> : <h1>Publicaciones Servidor1B</h1>
            }
            <span style={{marginLeft: '10%'}} />
            <button className="btn btn-primary" type="button" onClick={this.handleClick}>Cambiar servidor</button>
            <span style={{marginLeft: '10%'}} />
            cantidad:
            {
              this.state.users.length
            }
            <table className="table">
              <thead>
                <tr>
                  <th scope="col">Autor</th>
                  <th scope="col">Nota</th>
                </tr>
              </thead>
              <tbody>
                {this.state.users.map(function (user, index) {
                  return <tr key={index}><td>{user.autor}</td><td>{user.nota}</td></tr>
                })}


              </tbody>
            </table>
          </div>
        </div>
      </section>
    );
  }
}

export default Publicaciones;