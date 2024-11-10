import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../style.css';

function LoginPage() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLogin = () => {

        const isValidLogin = username === 'correctUsername' && password === 'correctPassword';

        if (isValidLogin) {
            navigate('/connection');
        } else {
            alert('Invalid username or password.');
        }
    };

    return (
        <div className="container">
            <h1>Login</h1>
            <input
                type="text"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
            />
            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <button onClick={handleLogin}>Login</button>
        </div>
    );
}

export default LoginPage;