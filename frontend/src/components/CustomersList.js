import React, { useEffect, useState } from 'react';
import '../App.css';

function CustomersList() {
  const [customers, setCustomers] = useState(null); // Инициализируем как null
  const [error, setError] = useState(null); // Состояние для ошибок

  useEffect(() => {
    fetch('http://localhost:8080/api/customers')
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => setCustomers(data))
      .catch(error => {
        console.error('Error fetching customers:', error);
        setError(error.message); // Устанавливаем сообщение об ошибке
        setCustomers([]);
      });
  }, []);

  // Если произошла ошибка, показываем сообщение
  if (error) {
    return <div>Error: {error}</div>;
  }

  // Если данные еще не загружены, показываем сообщение
  if (customers === null) {
    return <div>Loading customers...</div>;
  }

  // Если данные загружены, но массив пустой, показываем сообщение
  if (customers.length === 0) {
    return <div>No customers found.</div>;
  }

  return (
    <div className="customers-list">
      <h2>Customers</h2>
      <div className="list">
        {customers.map(customer => (
          <div key={customer.id} className="customer-card">
            <h3>{customer.first_name} {customer.last_name}</h3>
            <p>Email: {customer.email}</p>
            <p>Phone: {customer.phone_number}</p>
            <p>Loyalty Points: {customer.loyalty_points}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default CustomersList;