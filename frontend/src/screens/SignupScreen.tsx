import React, { SyntheticEvent } from 'react';
import { useState} from 'react';
import {Form, Button} from 'react-bootstrap'
import { RouteComponentProps } from 'react-router-dom';
import FormContainer from '../components/FormContainer'

interface Props {
  history: RouteComponentProps['history']
}

const SignupScreen = ({history}: Props) => {

  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword ] = useState('')
  const submitHandler = async(e: SyntheticEvent) =>{
    e.preventDefault()
    // Interact with the backend using fetch

    await fetch('http://localhost:8881/api/register', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({
        first_name: firstName,
        last_name: lastName,
        email,
        password,
      }),
    })
    history.push('/login')
  }
  return(
    <FormContainer>
    <h1>Sign Up</h1>
  <Form onSubmit={submitHandler}>
<Form.Group controlId="firstName" className='mb-3 my-3'>
  <Form.Label>First Name</Form.Label>
  <Form.Control type="firstName" placeholder="Enter your First Name" 
    value={firstName}
    onChange={e => setFirstName(e.target.value)}
  />
</Form.Group>

<Form.Group controlId="lastName" className='mb-3 my-3'>
  <Form.Label>Last Name</Form.Label>
  <Form.Control type="lastName" placeholder="Enter your Last Name" 
    value={lastName}
    onChange={e => setLastName(e.target.value)}
  />
</Form.Group>


<Form.Group controlId="email" className='mb-3 my-3'>
  <Form.Label>Email address</Form.Label>
  <Form.Control type="email" placeholder="Enter your email" 
    value={email}
    onChange={e => setEmail(e.target.value)}
  />
</Form.Group>


<Form.Group  controlId="password" className='mb-3 my-3'>
  <Form.Label>Password</Form.Label>
  <Form.Control type="password" placeholder="Password" 
    value={password}
    onChange={e => setPassword(e.target.value)}
  />
</Form.Group>

<Button variant="primary" type="submit" className='mb-3 my-3'>
  Submit
</Button>
</Form>
</FormContainer>
  )
}
export default SignupScreen;