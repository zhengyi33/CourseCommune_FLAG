import React, { useState } from "react";
import Main from "./Main";
import { TOKEN_KEY } from "../constants";
import "../styles/App.css";
import { LogoutOutlined } from '@ant-design/icons';

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(
        localStorage.getItem(TOKEN_KEY) ? true : false
    );

    const logout = () => {
        console.log("log out");
        localStorage.removeItem(TOKEN_KEY);
        setIsLoggedIn(false);
    };

    const loggedIn = (token) => {
        if (token) {
            localStorage.setItem(TOKEN_KEY, token);
            console.log(TOKEN_KEY)
            setIsLoggedIn(true);
        }
    };
    return (
        <div className="App">

            <LogoutOutlined className="logout" onClick={logout} />

            <Main isLoggedIn={isLoggedIn} handleLoggedIn={loggedIn} />
        </div>
    );
}

export default App;
