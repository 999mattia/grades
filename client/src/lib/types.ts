export interface User {
	id: number;
	name: string;
	modules: Module[];
}

export interface Module {
	id: number;
	name: string;
	semester: number;
	grades: Grade[];
}

export interface Grade {
	id: number;
	name: string;
	grade: number;
}
