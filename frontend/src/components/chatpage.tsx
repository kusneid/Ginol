import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import '../style.css';

interface Message {
    id: number;
    username: string;
    text: string;
    time: string;
}

const ChatPage: React.FC = () => {
    const [messages, setMessages] = useState<Message[]>([]);
    const [inputText, setInputText] = useState('');
    const location = useLocation();
    const { username, friend } = location.state as { username: string; friend: string };

    useEffect(() => {
        fetchMessages();
    }, []);

    const fetchMessages = async () => {
        try {
            const res = await fetch('http://localhost:8080/api/messages');
            const data = await res.json();
            setMessages(data);
        } catch (err) {
            console.error('Error fetching messages:', err);
        }
    };

    const handleSendMessage = async () => {
        if (inputText.trim()) {
            const newMessage = {
                username,
                text: inputText,
                time: new Date().toISOString(),
            };

            try {
                await fetch('http://localhost:5000/api/messages', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(newMessage),
                });

                await fetchMessages();
                setInputText('');
            } catch (err) {
                console.error('Error sending message:', err);
            }
        }
    };

    return (
        <div className="container">
            <h1 className="static-title">Chat with {friend}</h1>
            <div className="chat-box">
                {messages.map((msg) => (
                    <div key={msg.id} className={`message ${msg.username === username ? 'sent' : 'received'}`}>
                        <strong>{msg.username}</strong>: {msg.text}
                        <span className="time">{new Date(msg.time).toLocaleTimeString().slice(0, 5)}</span>
                    </div>
                ))}
            </div>
            <div className="input-container">
                <input
                    type="text"
                    placeholder="Write a message..."
                    value={inputText}
                    onChange={(e) => setInputText(e.target.value)}
                    onKeyDown={(e) => e.key === 'Enter' && handleSendMessage()}
                />
                <button onClick={handleSendMessage}>Send</button>
            </div>
        </div>
    );
};

export default ChatPage;