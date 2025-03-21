import React from 'react';
import { Routes, Route, Link } from 'react-router-dom'; // Импортируем только Routes, Route и Link
import Home from './components/Home';
import HotelList from './components/HotelList';
import RoomsList from './components/RoomsList';
import CustomersList from './components/CustomersList';
import BookingsList from './components/BookingsList';
import ReviewsList from './components/ReviewsList';
import './App.css';

function App() {
  return (
    <div className="App">
      <nav>
        <Link to="/">Home</Link>
        <Link to="/hotels">Hotels</Link>
        <Link to="/rooms">Rooms</Link>
        <Link to="/customers">Customers</Link>
        <Link to="/bookings">Bookings</Link>
        <Link to="/reviews">Reviews</Link>
      </nav>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/home" element={<Home />} />
        <Route path="/hotels" element={<HotelList />} />
        <Route path="/rooms" element={<RoomsList />} />
        <Route path="/customers" element={<CustomersList />} />
        <Route path="/bookings" element={<BookingsList />} />
        <Route path="/reviews" element={<ReviewsList />} />
      </Routes>
    </div>
  );
}

export default App;