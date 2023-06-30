export interface LoginData {
	name: string;
	password: string;
}

export interface DecodedToken {
	sub: string;
	exp: number;
}

export interface UserData {
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

export interface CreateModuleModel {
	name: string;
	semester: number | null;
}
