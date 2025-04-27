import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class StatisticsService {
  private apiUrl = 'http://localhost:1488/api/statistics';
  constructor(private http: HttpClient) { }

  getTopVisitedCountries(): Promise<any> {
    return firstValueFrom(this.http.get<any>(`${this.apiUrl}/top-visited`, { withCredentials: true }));
  }

  getTopWishlistCountries(): Promise<any> {
    return firstValueFrom(this.http.get<any>(`${this.apiUrl}/top-wish-list`, { withCredentials: true }));
  }
}
