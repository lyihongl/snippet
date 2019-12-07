import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Navbar from './NavigationBar/navbar.js'
//import App from './App';
//import * as serviceWorker from './serviceWorker';

//ReactDOM.render(<Test />, document.getElementById('root'));
//console.log(process.env.REACT_APP_API_KEY)
/*
class Test extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            //data: null,
            Name: null,
            Hobbies: null,
        };
        const API_KEY = process.env.REACT_APP_API_KEY;
        //console.log(API_KEY)
    }
    componentDidMount(){
        fetch("/test")
        .then(results => results.json()
        .then(data => this.setState({Name: data.Name, Hobbies: data.Hobbies})))
        console.log(this.state.Name)
    }
    render(){
        //this.test()
        return <h1>{this.state.Name}</h1>;
    }
}*/
class App extends React.Component{
    constructor(props) {
        super(props);
    }
    componentDidMount(){
    }
    render(){
        return (
            <div>
                <div class="navbar">
                    <Navbar />
                </div>
            </div>
        );
    }
}
ReactDOM.render(
    <App />,
    document.getElementById('root')
);
// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
//serviceWorker.unregister();
