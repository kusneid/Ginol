import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../style.css';

const Login: React.FC = () => {
    const [nickname, setNickname] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        const response = await fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ nickname, password }),
        });

        const data = await response.json();

        if (data.success) {

            navigate('/connection', { state: { nickname: data.nickname } });
        } else {

            alert('Login failed');
        }
    };

    return (
        <div className="form-container">
            <h1 className="form-title">Login</h1>
            <form onSubmit={handleLogin}>
                <input
                    type="text"
                    placeholder="Nickname"
                    value={nickname}
                    onChange={(e) => setNickname(e.target.value)}
                    className="form-input"
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="form-input"
                />
                <button type="submit" className="form-button">Login</button>
            </form>
        </div>
    );
};

export default Login;