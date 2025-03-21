import React from 'react';
import '../App.css';

function CancellationForm() {
  return (
    <div className="cancellation-form">
      <h2>Cancel Booking</h2>
      <form>
        <input type="text" placeholder="Booking ID" />
        <input type="email" placeholder="Your Email" />
        <button type="submit">Cancel Booking</button>
      </form>
    </div>
  );
}

export default CancellationForm;