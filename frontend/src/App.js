import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Home from './routes/home'
import Login from './routes/login'
import Register from './routes/register'
import EditAccount from "./routes/editAccount";
import ChangePassword from "./routes/changePassword";
import './App.css'
import { ChakraProvider } from '@chakra-ui/react'
import { useEffect, useState } from 'react'

function App() {
    const [loggedIn, setLoggedIn] = useState(false)
    const [username, setUsername] = useState("")

    console.log("App.js: loggedIn: ", loggedIn)

    return (
        <div className="App">
            <ChakraProvider>
                <BrowserRouter>
                    <Routes>
                        <Route path="/" element={<Home username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
                        <Route path="/login" element={<Login setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
                        <Route path="/register" element={<Register setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
                        <Route path="/account" element={<EditAccount
                            username={username}
                            loggedIn={loggedIn}
                            setLoggedIn={setLoggedIn}
                            setUsername={setUsername}
                        />} />
                        <Route path={"/password"} element={<ChangePassword
                            username={username}
                            loggedIn={loggedIn}
                            setLoggedIn={setLoggedIn}
                            setUsername={setUsername}
                        />} />
                    </Routes>
                </BrowserRouter>
            </ChakraProvider>
        </div>
    )
}

export default App