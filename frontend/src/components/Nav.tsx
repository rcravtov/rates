import { Component } from 'solid-js';

const Nav: Component = () => {

    const testClick = () => {
        console.log("test")
    }

    return (
        <nav class="navbar bg-body-tertiary">
            <div class="container-fluid">
               <a class="navbar-brand">Rates</a>
            </div>
        </nav>
    )
}

export default Nav;