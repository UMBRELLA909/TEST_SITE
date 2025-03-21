import React from 'react';
import '../App.css';

function BookingForm() {
  return (
    <div className="booking-form">
      <h2>Book a Room</h2>
      <form>
        <input type="text" placeholder="Your Name" />
        <input type="email" placeholder="Your Email" />
        <input type="date" placeholder="Check-in Date" />
        <input type="date" placeholder="Check-out Date" />
        <button type="submit">Book Now</button>
      </form>
    </div>
  );
}

export default BookingForm;