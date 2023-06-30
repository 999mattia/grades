import { Component } from '@angular/core';
import { UserService } from '../../shared/user.service';
import { Module, UserData } from '../../shared/models';
import { CalculateService } from '../../shared/calculate.service';
import { ModuleService } from '../../shared/module.service';
import { Router } from '@angular/router';

@Component({
	selector: 'app-home',
	templateUrl: './home.component.html',
})
export class HomeComponent {
	constructor(
		private userService: UserService,
		private calculateService: CalculateService,
		private moduleService: ModuleService,
		private router: Router
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
		});
	}

	calculateAverage(module: Module): number {
		return this.calculateService.calculateAverageForModule(module);
	}

	handleCreateModule() {
		this.router.navigate(['module/create']);
	}

	handleDeleteModule(id: number) {
		this.moduleService.deleteModule(id).subscribe(() => {
			this.fetchUserData();
		});
	}

	handleCreateGrade(moduleId: number) {
		this.router.navigate([`module/${moduleId}/grade/create`]);
	}
}
