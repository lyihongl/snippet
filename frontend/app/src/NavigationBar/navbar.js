import React from 'react';
import ReactDOM from 'react-dom';
import "./navbar.css"

class Navbar extends React.Component {
    constructor(props) {
        super(props)
    }
    componentDidMount() {

    }
    render() {
        return (
            <div>
                <nav class="container main-nav">
                    <div class="nav-item nav-1">
                        <a href="/projects" class="navLinks">Projects</a>
                    </div>
                    <div class="nav-item nav-2">
                        <a href="/" class="navLinks">Yihong Liu</a>
                    </div>
                    <div class="nav-item nav-3">
                        <a href="/about" class="navLinks">About Me</a>
                    </div>
                </nav>
            </div>
        );
    }
}

export default Navbar;