import React from 'react';
import { useNavigate } from 'react-router-dom';
import './styles.css';

const HomePage: React.FC = () => {
    const navigate = useNavigate();

    return (
        <div className="container">
            <h1>Welcome to Ginol</h1>
            <button onClick={() => navigate('/login')}>Login</button>
            <button onClick={() => navigate('/register')}>Register</button>
        </div>
    );
};

export default HomePage;