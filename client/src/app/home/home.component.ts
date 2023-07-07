import { Component } from '@angular/core';
import { UserService } from '../../shared/user.service';
import { Module, UserData } from '../../shared/models';
import { CalculateService } from '../../shared/calculate.service';
import { ModuleService } from '../../shared/module.service';
import { Router } from '@angular/router';
import { FormControl } from '@angular/forms';

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
	) {
		this.semester.valueChanges.subscribe((value) => {
			if (value == undefined) {
				this.modulesToShow = this.data.modules;
			} else {
				this.modulesToShow = this.data.modules.filter(
					(module) => module.semester == +value
				);
			}
		});
	}

	data: UserData = {
		id: 0,
		name: '',
		modules: [],
	};

	semester = new FormControl('');

	modulesToShow: Module[] = [];

	ngOnInit(): void {
		this.fetchUserData();
		this.modulesToShow = this.data.modules;
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
