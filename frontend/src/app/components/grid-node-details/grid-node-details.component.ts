import {
  Component,
  EventEmitter,
  Inject,
  Input,
  OnChanges,
  OnInit,
  Output,
  SimpleChanges,
} from '@angular/core';
import { Observable, last } from 'rxjs';
import { GridNode } from 'src/app/pb/grid_pb';
import { GridService } from 'src/app/services/grid/grid.service';

const CONSTANTS = {
  Header: 'Details',
};

@Component({
  selector: 'grid-node-details',
  templateUrl: './grid-node-details.component.html',
  styleUrls: ['./grid-node-details.component.scss'],
  providers: [{ provide: 'CONSTANTS', useValue: CONSTANTS }],
})
export class GridNodeDetailsComponent implements OnChanges {
  @Input() nodeId?: number;
  @Output() onSave = new EventEmitter<GridNode | null>();
  @Output() onCancel = new EventEmitter();

  node: GridNode | null = null;

  constructor(
    private gridService: GridService,
    @Inject('CONSTANTS') public Constants: any
  ) {}

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['nodeId']) {
      this.gridService.getNode(this.nodeId).subscribe((value) => {
        this.node = value ?? null;
      });
    } else {
      this.node = null;
    }
  }

  public handleSave() {
    this.onSave.emit(this.node);
  }

  public handleCancel(event: Event) {
    this.onCancel.emit();
    event.preventDefault();
  }
}
