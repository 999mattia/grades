import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { CreateModuleModel } from 'src/shared/models';
import { ModuleService } from 'src/shared/module.service';

@Component({
	selector: 'app-create-module',
	templateUrl: './create-module.component.html',
})
export class CreateModuleComponent {
	constructor(private moduleService: ModuleService, private router: Router) {}
	module: CreateModuleModel = {
		name: '',
		semester: null,
	};

	onSubmit() {
		this.moduleService.createModule(this.module).subscribe((res: any) => {
			this.router.navigate(['']);
		});
	}
}
