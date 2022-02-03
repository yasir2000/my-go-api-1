import React from 'react';
interface Props{
  firstName:string
}

export const HomeScreen = ({firstName}: Props) => {
  return (
    firstName ? <h1> Welcome {firstName}!</h1>:<h1>Welcome to the home page</h1>
  )
}

export default HomeScreen;