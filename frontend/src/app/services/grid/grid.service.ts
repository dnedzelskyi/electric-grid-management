import { Injectable } from '@angular/core';
import { GridServiceClient } from 'src/app/pb/grid_pb_service';
import { API } from '../utils/constants';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import {
  GetNodeRequest,
  GridNode,
  NodeLink,
  UpdateNodeRequest,
} from 'src/app/pb/grid_pb';
import { TApiGrid } from 'src/app/types/grid-types';
import { Observable, from, of } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class GridService {
  private client: GridServiceClient;

  constructor() {
    this.client = new GridServiceClient(API.Uri);
  }

  public getGrid(): Observable<TApiGrid> {
    const getGridPromise: Promise<TApiGrid> = new Promise((resolve) => {
      this.client.getGrid(new google_protobuf_empty_pb.Empty(), (err, res) =>
        resolve({
          nodes: res?.getNodesList() ?? [],
          links: res?.getLinksList() ?? [],
        })
      );
    });

    return from(getGridPromise);
  }

  public getNode(id?: number): Observable<GridNode | null> {
    if (!id) {
      return of(null);
    }

    const getNodePromise: Promise<GridNode | null> = new Promise((resolve) => {
      const req = new GetNodeRequest();
      req.setId(id);
      this.client.getNode(req, (err, res) => resolve(res?.getNode() ?? null));
    });

    return from(getNodePromise);
  }

  public updateNode(node: GridNode | null): Observable<number | undefined> {
    if (!node) {
      return of(undefined);
    }

    const updateNodePromise: Promise<number | undefined> = new Promise(
      (resolve) => {
        const req = new UpdateNodeRequest();
        req.setNode(node);
        this.client.updateNode(req, (err, res) => resolve(res?.getId()));
      }
    );

    return from(updateNodePromise);
  }
}
