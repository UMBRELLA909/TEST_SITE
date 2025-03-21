import React, { useEffect, useState } from 'react';
import '../App.css';

function HotelList() {
  const [hotels, setHotels] = useState(null); // Инициализируем как null
  const [error, setError] = useState(null); // Состояние для ошибок

  useEffect(() => {
    fetch('http://localhost:8080/api/hotels')
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => setHotels(data))
      .catch(error => {
        console.error('Error fetching hotels:', error);
        setError(error.message); // Устанавливаем сообщение об ошибке
        setHotels([]);
      });
  }, []);

  // Если произошла ошибка, показываем сообщение
  if (error) {
    return <div>Error: {error}</div>;
  }

  // Если данные еще не загружены, показываем сообщение
  if (hotels === null) {
    return <div>Loading hotels...</div>;
  }

  // Если данные загружены, но массив пустой, показываем сообщение
  if (hotels.length === 0) {
    return <div>No hotels found.</div>;
  }

  return (
    <div className="hotel-list">
      <h2>Hotels</h2>
      <div className="list">
        {hotels.map(hotel => (
          <div key={hotel.id} className="hotel-card">
            <h3>{hotel.name}</h3>
            <p>📍 {hotel.city}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default HotelList;