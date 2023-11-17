import React from 'react'
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";
import List from './tasks/List';



export default function App() {

  return (
    <Router>
      <Routes>
        <Route index element={<List />} />
        
      </Routes>
    </Router>
  );
}

