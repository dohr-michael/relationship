import { ApiService, UniversesService } from './services';
import { MaterialModule } from './modules';

export * from './models';
export const services = [ ApiService, UniversesService ];
export const modules = [ MaterialModule ];
