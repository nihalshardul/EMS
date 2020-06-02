import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DeletegroupComponent } from './deletegroup.component';

describe('DeletegroupComponent', () => {
  let component: DeletegroupComponent;
  let fixture: ComponentFixture<DeletegroupComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeletegroupComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeletegroupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
