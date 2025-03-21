import React, { useEffect, useState } from 'react';
import '../App.css';

function BookingsList() {
  const [bookings, setBookings] = useState(null); // Инициализируем как null
  const [error, setError] = useState(null); // Состояние для ошибок

  useEffect(() => {
    fetch('http://localhost:8080/api/bookings')
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => setBookings(data))
      .catch(error => {
        console.error('Error fetching bookings:', error);
        setError(error.message); // Устанавливаем сообщение об ошибке
        setBookings([]);
      });
  }, []);

  // Если произошла ошибка, показываем сообщение
  if (error) {
    return <div>Error: {error}</div>;
  }

  // Если данные еще не загружены, показываем сообщение
  if (bookings === null) {
    return <div>Loading bookings...</div>;
  }

  // Если данные загружены, но массив пустой, показываем сообщение
  if (bookings.length === 0) {
    return <div>No bookings found.</div>;
  }

  return (
    <div className="bookings-list">
      <h2>Bookings</h2>
      <div className="list">
        {bookings.map(booking => (
          <div key={booking.id} className="booking-card">
            <h3>Booking ID: {booking.id}</h3>
            <p>Hotel ID: {booking.hotel_id}</p>
            <p>Room ID: {booking.room_id}</p>
            <p>Customer ID: {booking.customer_id}</p>
            <p>Check-in: {new Date(booking.check_in_date).toLocaleDateString()}</p>
            <p>Check-out: {new Date(booking.check_out_date).toLocaleDateString()}</p>
            <p>Total Price: ${booking.total_price}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default BookingsList;