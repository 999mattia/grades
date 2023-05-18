import type { Component } from "solid-js";
import { Routes, Route, A } from "@solidjs/router";
import Login from "./Login";
import Home from "./Home";

import logo from "./logo.svg";
import styles from "./App.module.css";

const App: Component = () => {
	return (
		<>
			<Routes>
				<Route
					path="/"
					component={Home}
				/>
				<Route
					path="/login"
					component={Login}
				/>
			</Routes>
		</>
	);
};

export default App;
