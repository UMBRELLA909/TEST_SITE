import React, { useEffect, useState } from 'react';
import '../App.css';

function RoomsList() {
  const [rooms, setRooms] = useState(null); // Инициализируем как null
  const [error, setError] = useState(null); // Состояние для ошибок

  useEffect(() => {
    fetch('http://localhost:8080/api/rooms')
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => setRooms(data))
      .catch(error => {
        console.error('Error fetching rooms:', error);
        setError(error.message); // Устанавливаем сообщение об ошибке
        setRooms([]);
      });
  }, []);

  // Если произошла ошибка, показываем сообщение
  if (error) {
    return <div>Error: {error}</div>;
  }

  // Если данные еще не загружены, показываем сообщение
  if (rooms === null) {
    return <div>Loading rooms...</div>;
  }

  // Если данные загружены, но массив пустой, показываем сообщение
  if (rooms.length === 0) {
    return <div>No rooms found.</div>;
  }

  return (
    <div className="rooms-list">
      <h2>Rooms</h2>
      <div className="list">
        {rooms.map(room => (
          <div key={room.id} className="room-card">
            <h3>{room.category}</h3>
            <p>Hotel ID: {room.hotel_id}</p>
            <p>Capacity: {room.capacity}</p>
            <p>Price: ${room.price}</p>
            <p>Available: {room.available ? 'Yes' : 'No'}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default RoomsList;