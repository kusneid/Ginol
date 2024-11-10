import React, { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import '../style.css';

function ConnectionPage() {
    const [nickname, setNickname] = useState('');
    const navigate = useNavigate();
    const location = useLocation();
    const { nickname: loggedUserNickname } = (location.state as { nickname: string }) || {};

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setNickname(e.target.value);
    };

    const handleEnterChat = async () => {
        if (nickname.trim()) {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    alert("User is not authenticated");
                    navigate('/login');
                    return;
                }

                const response = await fetch('/api/check-nickname', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                    body: JSON.stringify({
                        loggedUser: loggedUserNickname,
                        friendNickname: nickname
                    })
                });

                if (!response.ok) {
                    console.error(`Error: Received status ${response.status}`);
                    alert("An error occurred. Please try again.");
                    return;
                }

                const data = await response.json();
                if (data.exists) {
                    navigate(`/chat/${nickname}`, { state: { username: loggedUserNickname, friend: nickname } });
                } else {
                    alert("Nickname not found in the database.");
                }
            } catch (error) {
                console.error("Error checking nickname:", error);
                alert("An error occurred. Please try again.");
            }
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