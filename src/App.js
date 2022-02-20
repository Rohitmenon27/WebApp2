import React, {Component, Fragment } from 'react';
import {BrowserRouter as Router, Switch, Route, Link, useParams} from 'react-router-dom';
import Movies from './components/movies';
import Home from './components/home';
import Admin from './components/admin';
import OneMovie from './components/OneMovie';
import Genres from './components/genres';

export default function App() {
  return (
    <Router>
    <div className="container">
      <div className="row">
        <h1 className="mt-3">
          List of movies to watch 
        </h1>
        <hr className="mb-3"></hr>
      </div>

      <div className= "row">
        <div className="col-md-2">
          <nav>
            <ul className= "list-group">
              <li className="list-group-items">
                <Link to="/">Home</Link>
              </li>
              <li className="list-group-items">
              <Link to="/movies">Movies</Link>
              </li>
              <li className="list-group-items">
              <Link to="/admin">Admin</Link>
              </li>
              <li className="list-group-items">
              <Link to="/genres">Genres</Link>
              </li>
            </ul>
          </nav>  
      </div>
      <div className="col-md-10">
        <Switch>
         <Route path="/movies/:id" component ={OneMovie} />
         <Route path="/movies">
            <Movies />
        </Route>
        <Route exact path = "/genres">
          <Genres />
         </Route>
          <Route path ="/admin">
            <Admin />
          </Route>x
         <Route path ="/">
            <Home/>
          </Route>
        </Switch>
      </div>
     </div>
    </div>
    </Router>
  );
}



