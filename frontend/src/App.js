import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Home from './components/home'
import Login from './components/login'
import Register from './components/register'
import './App.css'
import { ChakraProvider } from '@chakra-ui/react'
import { useEffect, useState } from 'react'

function App() {
    const [loggedIn, setLoggedIn] = useState(false)
    const [username, setUsername] = useState("")

    return (
        <div className="App">
            <ChakraProvider>
                <BrowserRouter>
                    <Routes>
                        <Route path="/" element={<Home username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
                        <Route path="/login" element={<Login setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
                        <Route path="/register" element={<Register setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
                    </Routes>
                </BrowserRouter>
            </ChakraProvider>
        </div>
    )
}

export default App