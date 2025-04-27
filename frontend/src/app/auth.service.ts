import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable,BehaviorSubject } from 'rxjs';

export interface UserProfile {
  id: number;
  username: string;
  email: string;
  created_at: string;
}

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private apiUrl = 'http://localhost:1488/api';
  private isLoggedInSubject = new BehaviorSubject<boolean>(false);
  private userDataSubject = new BehaviorSubject<UserProfile | null>(null);
  constructor(private http: HttpClient) {}

  login(email: string, password: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/auth/login`, { email, password }, { withCredentials: true });
  }

  register(username: string, email: string, password: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/auth/register`, { username, email, password }, { withCredentials: true });
  }

  profile(): Observable<UserProfile> {
    return this.http.get<UserProfile>(`${this.apiUrl}/users/profile`, { withCredentials: true });
  }

  logout(): Observable<any> {
    return this.http.post(`${this.apiUrl}/auth/logout`, {}, { withCredentials: true });
  }

  checkLoginStatus(): void {
    this.profile().subscribe({
      next: (response) => {
        this.isLoggedInSubject.next(true);
        this.userDataSubject.next(response);
      },
      error: () => {
        this.isLoggedInSubject.next(false);
        this.userDataSubject.next(null);
      }
    });
  }

  get isLoggedIn$(): Observable<boolean> {
    return this.isLoggedInSubject.asObservable();
  }

  get userData$(): Observable<UserProfile | null> {
    return this.userDataSubject.asObservable();
  }

}