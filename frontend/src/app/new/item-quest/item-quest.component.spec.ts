import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ItemQuestComponent } from './item-quest.component';

describe('ItemQuestComponent', () => {
  let component: ItemQuestComponent;
  let fixture: ComponentFixture<ItemQuestComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ItemQuestComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ItemQuestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
