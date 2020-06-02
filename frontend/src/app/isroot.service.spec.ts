import { TestBed } from '@angular/core/testing';

import { IsrootService } from './isroot.service';

describe('IsrootService', () => {
  let service: IsrootService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IsrootService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
