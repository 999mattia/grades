import type { Component } from "solid-js";
import { Routes, Route, A } from "@solidjs/router";

const Home: Component = () => {
	return (
		<>
			Home
			<A href="/login">Login</A>
		</>
	);
};

export default Home;
