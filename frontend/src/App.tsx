import React from 'react';
import {useEffect, useState} from 'react';
import './App.css';
import Header from './components/Header'
import Footer from './components/Footer'
import {Container} from 'react-bootstrap'
import {BrowserRouter as Router, Route } from 'react-router-dom'
import HomeScreen from './screens/HomeScreen'
import LoginScreen from './screens/LoginScreen'
import SignupScreen from './screens/SignupScreen'

function App() {
  const [firstName, setFirstName] = useState('')
  useEffect(() => {
    (
    async () => {
      const response = await fetch('http://localhost:8881:/api/user', {
        headers: {'Content-Type': 'application/json'},
        credentials: 'include',      
      })
      const data = await response.json()
      setFirstName(data.first_name)
    
    }
    )()
  })
  return (
    <Router>
    <Header firstName={firstName}/>
    <main>
      <Container>
            <Route path='/' component={() => <HomeScreen firstName={firstName}/>}/>
            <Route path='/signup' component={SignupScreen} />
            <Route path='/login' component={LoginScreen} />
      </Container>
    </main>
     <Footer/>
    </Router>
  )
}

export default App;
