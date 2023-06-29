import { Component } from '@angular/core';
import { UserService } from '../shared/user.service';
import { UserData } from '../shared/models';

@Component({
	selector: 'app-home',
	templateUrl: './home.component.html',
	styleUrls: ['./home.component.css'],
})
export class HomeComponent {
	constructor(private userService: UserService) {}

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
}
