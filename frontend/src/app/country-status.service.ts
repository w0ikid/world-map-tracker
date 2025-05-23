import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { environment } from '../environments/environment';
@Injectable({
  providedIn: 'root'
})
export class CountryStatusService {
  private apiUrl = environment.apiUrl;
  // private apiUrl = 'http://localhost:1488/api';
  constructor(private http: HttpClient) {}

  getCountryStatuses(): Observable<Record<string, 'visited' | 'wishlist'>> {
    return this.http.get<any[]>(`${this.apiUrl}/countries/`, { withCredentials: true }).pipe(
      map(statuses => {
        const statusMap: Record<string, 'visited' | 'wishlist'> = {};
        statuses.forEach(status => {
          if (status.status === 'visited' || status.status === 'wishlist') {
            statusMap[status.country_iso] = status.status;
          }
        });
        return statusMap;
      }),
      catchError(this.handleError)
    );
  }

  getVisitedPercentage(): Observable<{ visited_percentage: number }> {
    return this.http.get<{ visited_percentage: number }>(`${this.apiUrl}/countries/visited-percentage`, { withCredentials: true });
  }

  getVisitedCount(): Observable<{ visited_count: number }> {
    return this.http.get<{ visited_count: number }>(`${this.apiUrl}/countries/visited-count`, { withCredentials: true });
  }

  getWishlistCount(): Observable<{ wishlist_count: number }> {
    return this.http.get<{ wishlist_count: number }>(`${this.apiUrl}/countries/wish-list-count`, { withCredentials: true });
  }


  // Извините за путаницу, но это не нужно :> / Быстрая реализация треубует харкодинга

  getVisitedPercentageByUsername(username: string): Observable<{ visited_percentage: number }> {
    return this.http.get<{ visited_percentage: number }>(`${this.apiUrl}/countries/visited-percentage/${username}`, { withCredentials: true });
  }
  
  getVisitedCountByUsername(username: string): Observable<{ visited_count: number }> {
    return this.http.get<{ visited_count: number }>(`${this.apiUrl}/countries/visited-count/${username}`, { withCredentials: true });
  }
  
  getWishlistCountByUsername(username: string): Observable<{ wishlist_count: number }> {
    return this.http.get<{ wishlist_count: number }>(`${this.apiUrl}/countries/wish-list-count/${username}`, { withCredentials: true });
  }

  setCountryStatus(countryISO: string, status: 'visited' | 'wishlist' | 'none'): Observable<any> {
    const body = { country_iso: countryISO, status };
    if (status === 'none') {
      return this.http.delete(`${this.apiUrl}/countries/`, { body, withCredentials: true }).pipe(
        catchError(this.handleError)
      );
    } else {
      return this.http.post(`${this.apiUrl}/countries/`, body, { withCredentials: true }).pipe(
        catchError(this.handleError)
      );
    }
  }

  updateCountryStatus(countryISO: string, status: 'visited' | 'wishlist'): Observable<any> {
    const body = { country_iso: countryISO, status };
    return this.http.put(`${this.apiUrl}/countries/`, body, { withCredentials: true }).pipe(
      catchError(this.handleError)
    );
  }

  deleteCountryStatus(countryISO: string): Observable<any> {
    const body = { country_iso: countryISO };
    return this.http.delete(`${this.apiUrl}/countries/`, { body, withCredentials: true }).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: HttpErrorResponse) {
    let errorMessage = 'An unknown error occurred!';
    if (error.error instanceof ErrorEvent) {
      errorMessage = `Error: ${error.error.message}`;
    } else {
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.error?.error || error.message}`;
    }
    return throwError(() => new Error(errorMessage));
  }
}