import { Component, Inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Observable, filter, map } from 'rxjs';
import { GridNode } from 'src/app/pb/grid_pb';
import { GridService } from 'src/app/services/grid/grid.service';
import { TApiGrid, TGraphGrid } from 'src/app/types/grid-types';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss'],
})
export class HomePageComponent implements OnInit {
  selectedNodeId: number | undefined;
  grid$!: Observable<TGraphGrid>;

  constructor(
    private gridService: GridService,
    private route: ActivatedRoute,
    private router: Router
  ) {
    this.refresh();
  }

  ngOnInit(): void {
    this.route.params.subscribe((params) => {
      if ('id' in params) {
        this.selectedNodeId = parseInt(params['id']);
      }
    });
  }

  handleSave(node: GridNode | null) {
    this.gridService.updateNode(node).subscribe(() => {
      this.router.navigate(['/home']);
      this.refresh();
    });
  }

  handleCancel() {
    this.router.navigate(['/home']);
    this.refresh(true);
  }

  private handleFetchAll() {
    this.grid$ = this.gridService.getGrid().pipe(
      map<TApiGrid, TGraphGrid>((g) => ({
        nodes: g.nodes.map((n) => ({
          id: n.getId().toString(),
          label: n.getLabel(),
          data: { type: n.getType(), power: n.getPower() },
        })),
        links: g.links.map((l) => ({
          source: l.getSource().toString(),
          target: l.getTarget().toString(),
        })),
      }))
    );
  }

  private refresh(skipData: boolean = false) {
    !skipData && this.handleFetchAll();
  }
}
