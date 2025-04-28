import { Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { ProfileComponent } from './profile/profile.component';
import { RegisterComponent } from './register/register.component';
import { WorldMapComponent } from './world-map/world-map.component';
import { AboutComponent } from './about/about.component';
import { StatisticsComponent } from './statistics/statistics.component';
import { UserProfileComponent } from './user-profile/user-profile.component';
export const routes: Routes = [
    { path: '', redirectTo: '/login', pathMatch: 'full' },
    { path: 'login', component: LoginComponent},
    { path: 'register', component: RegisterComponent},
    { path: 'profile', component: ProfileComponent},
    { path: 'worldmap', component: WorldMapComponent},
    { path: 'about', component: AboutComponent},
    { path: 'statistics', component: StatisticsComponent},
    { path: 'user/:username', component: UserProfileComponent},
];

