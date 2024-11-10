import React, { useState, useEffect, useRef } from 'react';
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
    const { username, friend, userToken, friendToken } = location.state as {
        username: string;
        friend: string;
        userToken: string;
        friendToken: string;
    };
    const chatBoxRef = useRef<HTMLDivElement>(null);
    const [socket, setSocket] = useState<WebSocket | null>(null);

    useEffect(() => {
        const ws = new WebSocket(`ws://localhost:8080/ws?token=${userToken}`);
        setSocket(ws);

        ws.onmessage = (event) => {
            const message: Message = JSON.parse(event.data);
            if ([message.username, friend].includes(username) && [username, message.username].includes(friend)) {
                setMessages((prevMessages) => [...prevMessages, message]);
            }
        };

        ws.onclose = () => {
            console.log('WebSocket connection closed');
        };

        return () => {
            ws.close();
        };
    }, [userToken, username, friend]);

    useEffect(() => {
        chatBoxRef.current?.scrollTo({
            top: chatBoxRef.current.scrollHeight,
            behavior: 'smooth',
        });
    }, [messages]);

    const handleSendMessage = () => {
        if (inputText.trim() && socket) {
            const newMessage: Message = {
                id: Date.now(),
                username,
                text: inputText,
                time: new Date().toISOString(),
            };

            socket.send(
                JSON.stringify({
                    ...newMessage,
                    token: userToken,
                    recipientToken: friendToken,
                })
            );

            setMessages((prevMessages) => [...prevMessages, newMessage]);
            setInputText('');
        }
    };

    return (
        <div className="container">
            <h1 className="static-title">Chat with {friend}</h1>
            <div className="chat-box" ref={chatBoxRef}>
                {messages.map((msg, index) => (
                    <div key={index} className={`message ${msg.username === username ? 'sent' : 'received'}`}>
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
                <button onClick={handleSendMessage} disabled={!inputText.trim()}>
                    Send
                </button>
            </div>
        </div>
    );
};

export default ChatPage;