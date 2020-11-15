import React from 'react';
import './App.css';
import Navbar from './Components/Navbar'
import Metricas from './Components/Metricas'

function App() {
  return (
    <header>
      <Navbar />
      <div className="container">
        <Metricas />
      </div>

    </header>
  );
}

export default App;
