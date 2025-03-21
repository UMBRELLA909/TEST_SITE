import React, { useEffect, useState } from 'react';
import '../App.css';

function ReviewsList() {
  const [reviews, setReviews] = useState(null); // Инициализируем как null
  const [error, setError] = useState(null); // Состояние для ошибок

  useEffect(() => {
    fetch('http://localhost:8080/api/reviews')
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => setReviews(data))
      .catch(error => {
        console.error('Error fetching reviews:', error);
        setError(error.message); // Устанавливаем сообщение об ошибке
        setReviews([]);
      });
  }, []);

  // Если произошла ошибка, показываем сообщение
  if (error) {
    return <div>Error: {error}</div>;
  }

  // Если данные еще не загружены, показываем сообщение
  if (reviews === null) {
    return <div>Loading reviews...</div>;
  }

  // Если данные загружены, но массив пустой, показываем сообщение
  if (reviews.length === 0) {
    return <div>No reviews found.</div>;
  }

  return (
    <div className="reviews-list">
      <h2>Reviews</h2>
      <div className="list">
        {reviews.map(review => (
          <div key={review.id} className="review-card">
            <h3>Rating: {review.rating}/5</h3>
            <p>Hotel ID: {review.hotel_id}</p>
            <p>Customer ID: {review.customer_id}</p>
            <p>Comment: {review.comment}</p>
            <p>Date: {new Date(review.review_date).toLocaleDateString()}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default ReviewsList;