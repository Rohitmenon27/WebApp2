import React, {Component, Fragment} from 'react';
import {Link} from 'react-router-dom';

export default class Movies extends Component {
    state = {
         movies: [],
         isLoaded: false,
         error: null
        
    }; 

    componentDidMount() {

        fetch("http://localhost:5000/v1/movies")
          .then((response) => {
              console.log("Status code is", response.status)
              if (response.status !== "200") {
                  let err = Error;
                  err.message = "Invalid Response:" + response.status;
                  this.setState({error: err}); 
              }
              return response.json();
          })
          .then((json) => {
              this.setState({
                  movies: json.movies,
                  isLoaded: true,
              },
              (error) => {
                  this.setState( {
                      isLoaded: true,
                      error
                  });
              }
              
              );

          });

    }
    render() {
        const { movies, isLoaded, error} = this.state;
        if (error) {
            return <div>Error: {error.message}</div>
        }   else if (!isLoaded) {
            return <p>Wait Loading...</p>;
        }   else {return (
            <Fragment>
             <h2>Choose a Movie </h2>

             <ul>
                {movies.map((m) => (
                    <li key={m.id}>
                       <Link to={`/movies/${m.id}`} >{m.title}</Link>
                    </li>
                ))}
            </ul>
            </Fragment>
        );
    
    }
  }
}