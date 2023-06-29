import { Component } from '@angular/core';
import { UserService } from '../shared/user.service';
import { Module, UserData } from '../shared/models';
import { CalculateService } from '../shared/calculate.service';

@Component({
	selector: 'app-home',
	templateUrl: './home.component.html',
	styleUrls: ['./home.component.css'],
})
export class HomeComponent {
	constructor(
		private userService: UserService,
		private calculateService: CalculateService
	) {}

	data: UserData = {
		id: 0,
		name: '',
		modules: [],
	};

	ngOnInit(): void {
		this.fetchUserData();
	}

	fetchUserData() {
		this.userService.getUserData().subscribe((data: UserData) => {
			this.data = data;
			console.log(this.data.id);
			console.log(this.data.name);
			console.log(this.data.modules);
		});
	}

	calculateAverage(module: Module): number {
		return this.calculateService.calculateAverageForModule(module);
	}
}
