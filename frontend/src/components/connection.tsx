import React, { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import '../style.css';

function ConnectionPage() {
    const [nickname, setNickname] = useState('');
    const navigate = useNavigate();
    const location = useLocation();
    const { nickname: loggedUserNickname } = location.state as { nickname: string };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setNickname(e.target.value);
    };

    const handleEnterChat = () => {
        if (nickname.trim()) {
            navigate(`/chat/${nickname}`, { state: { loggedUserNickname } });
        } else {
            alert('Please enter a nickname.');
        }
    };

    return (
        <div className="container">
            <h1 style={{ marginBottom: '0', fontSize: '2rem' }}>Enter nickname</h1>
            <h1 style={{ marginTop: '0', fontSize: '2rem' }}>of your friend:</h1>
            <input
                type="text"
                placeholder="Type your friend's nickname..."
                value={nickname}
                onChange={handleInputChange}
            />
            <button onClick={handleEnterChat}>Enter chat</button>
        </div>
    );
}

export default ConnectionPage;