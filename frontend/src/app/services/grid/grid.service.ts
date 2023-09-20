import { Injectable } from '@angular/core';
import { GridServiceClient } from 'src/app/pb/grid_pb_service';
import { API } from '../utils/constants';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import { GridNode, NodeLink } from 'src/app/pb/grid_pb';
import { TApiGrid } from 'src/app/types/grid-types';

@Injectable({
  providedIn: 'root',
})
export class GridService {
  private client: GridServiceClient;

  constructor() {
    this.client = new GridServiceClient(API.Uri);
  }

  public async getGrid(): Promise<TApiGrid> {
    const getGridPromise: Promise<TApiGrid> = new Promise((resolve) => {
      this.client.getGrid(new google_protobuf_empty_pb.Empty(), (err, res) =>
        resolve({
          nodes: res?.getNodesList() ?? [],
          links: res?.getLinksList() ?? [],
        })
      );
    });

    return getGridPromise;
  }
}
