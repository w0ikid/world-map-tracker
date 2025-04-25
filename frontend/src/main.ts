import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';
import { provideRouter } from '@angular/router';
import { provideHttpClient, withFetch } from '@angular/common/http'; // Импортируем провайдеры HTTP
import { routes } from './app/app.routes';

bootstrapApplication(AppComponent, {
  providers: [  
    provideRouter(routes),
    provideHttpClient(withFetch()), // Предоставляем HttpClient глобально
  ],
}).catch((err) => console.error(err));