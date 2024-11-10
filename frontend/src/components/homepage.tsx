import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import '../style.css';

const HomePage: React.FC = () => {
    const navigate = useNavigate();
    const [animateTitle, setAnimateTitle] = useState(true);

    useEffect(() => {
        const timer = setTimeout(() => setAnimateTitle(false), 3000); // Анимация 3 секунды
        return () => clearTimeout(timer);
    }, []);

    return (
        <div className="container">
            <h1 className={animateTitle ? 'typing-animation' : 'static-title'}>Welcome to Ginol</h1>
            <button onClick={() => navigate('/login')}>Login</button>
            <button onClick={() => navigate('/register')}>Register</button>
        </div>
    );
};

export default HomePage;