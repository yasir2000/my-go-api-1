import React, { SyntheticEvent } from 'react';
import { useState} from 'react'
import {Form, Button} from 'react-bootstrap'
import { RouteComponentProps } from 'react-router-dom';
import FormContainer from '../components/FormContainer'

interface Props {
  history: RouteComponentProps['history']
}

export const LoginScreen = ({history}: Props) => {

  const [email, setEmail] = useState('')
  const [password, setPassword ] = useState('')
  const submitHandler =async(e: SyntheticEvent) =>{
      e.preventDefault()
      // Interact with the backend using fetch
      await fetch('http://localhost:8881/api/login', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      credentials: 'include',
      body: JSON.stringify({   
        email,
        password,
      }),
    })
    history.push('/')
  }
  return (
    <FormContainer>
      <h1>Login</h1>
    <Form onSubmit={submitHandler}>
  <Form.Group controlId="email" className='mb-3 my-3'>
    <Form.Label>Email address</Form.Label>
    <Form.Control type="email" placeholder="Enter your email" 
    value={email}
    onChange={e => setEmail(e.target.value)}/>
    
  </Form.Group>

  <Form.Group  controlId="password" className='mb-3 my-3'>
    <Form.Label>Password</Form.Label>
    <Form.Control type="password" placeholder="Password" 
    value={password}
    onChange={e => setPassword(e.target.value) }
    />
  </Form.Group>
  
  <Button variant="primary" type="submit" className='mb-3 my-3'>
    Submit
  </Button>
</Form>
</FormContainer>
  )
}
export default LoginScreen;