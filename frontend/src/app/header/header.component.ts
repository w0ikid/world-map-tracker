import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router, RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-header',
  imports: [CommonModule, RouterLink, FormsModule],
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  isLoggedIn: boolean = false;
  userData: any = null;
  isMenuOpen: boolean = false;
  searchUsername = '';
  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.authService.checkLoginStatus();
    
    this.authService.isLoggedIn$.subscribe(isLoggedIn => {
      this.isLoggedIn = isLoggedIn;
    });

    // Подписка на изменения данных пользователя
    this.authService.userData$.subscribe(userData => {
      this.userData = userData;
    });
  }

  logout(): void {
    this.authService.logout().subscribe({
      next: () => {
        this.isLoggedIn = false;
        this.userData = null;
        this.router.navigate(['/login']);
      },
      error: (error) => {
        console.error('Ошибка при выходе из системы:', error);
      }
    });
  }

  navigateToUser(): void {
    const username = this.searchUsername.trim();

    if (!username) {
      console.warn('Поиск пользователя: поле пустое.');
      return;
    }

    this.router.navigate(['/user', username]);

    this.searchUsername = '';
    this.toggleMenu();
  }

  toggleMenu(): void {
    this.isMenuOpen = !this.isMenuOpen;
  }
}
