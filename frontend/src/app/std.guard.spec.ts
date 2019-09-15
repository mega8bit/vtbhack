import { TestBed, async, inject } from '@angular/core/testing';

import { StdGuard } from './std.guard';

describe('StdGuard', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [StdGuard]
    });
  });

  it('should ...', inject([StdGuard], (guard: StdGuard) => {
    expect(guard).toBeTruthy();
  }));
});
