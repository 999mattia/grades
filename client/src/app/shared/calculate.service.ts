import { Injectable } from '@angular/core';
import { Module } from './models';

@Injectable({
	providedIn: 'root',
})
export class CalculateService {
	constructor() {}

	calculateAverageForModule(module: Module): number {
		let sum = 0;
		module.grades.forEach((grade) => {
			sum += grade.grade;
		});
		return sum / module.grades.length;
	}
}
