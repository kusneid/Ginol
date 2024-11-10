import React, { useState } from 'react';
import '../style.css';

function ConnectionPage() {
    const [nickname, setNickname] = useState('');

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setNickname(e.target.value);
    };

    const handleEnterChat = () => {
        if (nickname.trim()) {
            console.log(`Entering chat as ${nickname}`);
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
                placeholder="type your friend's nickname..."
                value={nickname}
                onChange={handleInputChange}
            />
            <button onClick={handleEnterChat}>Enter chat</button>
        </div>
    );
}

export default ConnectionPage;