import React, { useEffect, useState } from 'react';
import goServer from '../../api/go-server';

function Home() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    goServer.getUsers().then((response) => {
      setUsers(response.data);
    }).catch((err) => {
      console.log(err);
    });
  }, []);

  return(
    <div>
      <h2>Home</h2>
      <h3>Users list</h3>
      <ul>
        {users.map((user) => (
          <li key={user.ID}>{user.Username} - {user.Email}</li>
        ))}
      </ul>
    </div>
  );
}

export default Home;
