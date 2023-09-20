import { Component, Inject, Input } from '@angular/core';

const CONSTANTS = {
  Header: 'Details',
};

@Component({
  selector: 'grid-node-details',
  templateUrl: './grid-node-details.component.html',
  styleUrls: ['./grid-node-details.component.scss'],
  providers: [{ provide: 'CONSTANTS', useValue: CONSTANTS }],
})
export class GridNodeDetailsComponent {
  @Input() nodeId?: number;
  constructor(@Inject('CONSTANTS') public Constants: any) {}
}
